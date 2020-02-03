const fdk = require('@fnproject/fdk');

fdk.handle(function (input) {
  let l = 0
  let r = 0
  if (input.left) {
    l = Number(input.left)
  }
  if (input.right) {
    r = Number(input.right)
  }
  const result = l + r
  return { 'result': String(result) }
})
