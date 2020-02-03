require 'fdk'
require 'bigdecimal'

def subtract(context:, input:)
  left = input.respond_to?(:fetch) ? input.fetch('left') : "0"
  right = input.respond_to?(:fetch) ? input.fetch('right') : "0"
  l = BigDecimal(left)
  r = BigDecimal(right)
  { result: (l - r).to_s("F") }
end

FDK.handle(target: :subtract)
