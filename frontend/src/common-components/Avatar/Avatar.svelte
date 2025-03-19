<script lang="ts">
  import Icon from '@iconify/svelte'
  import AvatarLayout from './AvatarLayout.svelte'
  import { calcAvatarSize } from './helpers'

  type AvatarProps = {
    src: string | null
    alt: string
    size?: 'extra small' | 'small' | 'large' | '24' | string
    borderRadius?: 'full' | 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | '3xl' | 'none'
    mobileSize?: 'extra small' | 'small' | 'large' | '24' | string
    noSrcIcon?: string
    noSrcImage?: string
  }

  let {
    src,
    alt,
    size = 'small',
    borderRadius = 'full',
    mobileSize = size,
    noSrcIcon = 'hugeicons:user-circle-02',
    noSrcImage,
  }: AvatarProps = $props()

  const borderRadiusStyle = `rounded-${borderRadius}`

  const sizeClass = calcAvatarSize(size, 'lg:')
  const mobileSizeClass = calcAvatarSize(mobileSize, '')
</script>

<AvatarLayout>
  {#if src || noSrcImage}
    <img
      class="{borderRadiusStyle} object-center object-cover {sizeClass} {mobileSizeClass}  "
      src={src ? src : noSrcImage}
      {alt}
    />
  {:else}
    <Icon
      class="{borderRadiusStyle} text-cyan-500 size-14 {sizeClass} {mobileSizeClass} "
      icon={noSrcIcon}
    />
  {/if}
</AvatarLayout>
