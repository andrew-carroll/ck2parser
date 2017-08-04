require_relative './ck2hash.rb'
require 'ruby-progressbar'
class CK2Parser
  attr_reader :base_hash
  def initialize(filepath)
    @file = File.readlines(filepath).reverse.map(&:scrub)
    @base_hash = CK2Hash.new(nil, 1)
  end

  def parse!
    fail_unless_valid_save!
    prog_format = '%a %E |%w%i|'
    @progress = ProgressBar.create(total: @file.length, format: prog_format)
    scan_line!(@file.pop, @base_hash) until @file.empty?
  end

  private

  def fail_unless_valid_save!
    err_msg = 'Not a valid CK2 save file. Make sure file is uncompressed.'
    raise(ArgumentError, err_msg) unless @file.pop.match(/\ACK2txt\n/)
    @base_hash.open!
  end

  def scan_line!(line, hash)
    level = hash.level
    word = "[\w\d_]+"
    case
    when pattern(0, :empty_line).match(line)
      # do nothing
    when pattern(level, :open_hash).match(line)
      hash.open!
    when pattern(level, :close_hash).match(line)
      hash.close!
    when pattern(1, :close_hash).match(line)
      hash.close!
    when pattern(level, :new_hash).match(line)
      add_hash($1, level, hash)
    when pattern(level, :new_unnamed_hash).match(line)
      add_hash(nil, level, hash, :open)
    when pattern(level, :new_hash_same_line).match(line)
      add_hash($1, level, hash, :open)
    when pattern(level, :new_hash_same_line_no_indent).match(line)
      add_hash($1, level-1, hash, :open)
    when pattern(level, :property).match(line)
      add_property($1, $2, hash)
    else
      binding.pry
    end
    @progress.increment
  rescue => e
    binding.pry
    raise(e)
  end

  def pattern(level, pat)
    word = /[\w\d_\.-]+/
    case pat
    when :empty_line
      /#{tabs(level)}\n/
    when :new_hash 
      /#{tabs(level)}(#{word})=\n/
    when :new_unnamed_hash
      /#{tabs(level)}\{\n/
    when :new_hash_same_line
      /#{tabs(level)}(#{word})=\{\n/
    when :new_hash_same_line_no_indent
      /^(#{word})=\{\n/
    when :open_hash 
      /#{tabs(level-1)}\{\n/
    when :close_hash 
      /#{tabs(level-1)}\}\n/
    when :property 
      /#{tabs(level)}(#{word})=(.+)\n/
    end
  end

  def add_hash(key, level, hash, init_status=:new)
    new_hash = CK2Hash.new(key, level+1, init_status)
    scan_line!(@file.pop, new_hash) until new_hash.status == :closed
    hash.add_hash(new_hash)
    new_hash
  end

  def add_property(key, value, hash)
    hash.add_property(key, value)
  end

  def tabs(level)
    "^\t{#{level}}"
  end

end
