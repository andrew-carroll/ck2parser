class CK2Hash
  attr_reader :key, :level, :subhashes, :properties, :status
  def initialize(key, level, init_status=:new)
    @key = key
    @level = level
    @subhashes = Array.new
    @properties = Hash.new
    @status = init_status
  end

  def open!
    raise("Hash is already open!") if status == :open
    @status = :open
  end

  def close!
    fail_unless_open!
    @status = :closed
  end

  def add_property(key, value)
    fail_unless_open!
    @properties[key] = value
  end

  def add_hash(hash)
    fail_unless_open!
    @subhashes << hash
  end

  private

  def fail_unless_open!
    raise("Hash is not yet open!") if status == :new
    raise("Hash is already closed!") if status == :closed
  end
end
