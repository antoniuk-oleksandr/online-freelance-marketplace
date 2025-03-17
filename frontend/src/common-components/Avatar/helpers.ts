export const calcAvatarSize = (size: string, prefix: string): string => {
  let val
  switch (size) {
    case 'extra small':
      val = 'size-12'
      break
    case 'small':
      val = 'size-16'
      break
    case 'large':
      val = 'size-32'
      break
    default:
      val = `size-${size}`
      break
  }

  return prefix + val
}
