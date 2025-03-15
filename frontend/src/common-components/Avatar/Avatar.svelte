<script lang="ts">
  import Icon from '@iconify/svelte'
  import AvatarLayout from './AvatarLayout.svelte'

  type AvatarProps = {
    src: string | null
    alt: string
    size?: 'extra small' | 'small' | 'large' | '24' | string
  }

  const { src, alt, size = 'small' }: AvatarProps = $props()

  let sizeClass = $state('')
  switch (size) {
    case 'extra small':
      sizeClass = 'size-12'
      break
    case 'small':
      sizeClass = 'size-16'
      break
    case 'large':
      sizeClass = 'size-32'
      break
    case '24':
      sizeClass = 'size-24'
      break
    default:
      sizeClass = `size-${size}`
      break
  }
</script>

<AvatarLayout>
  {#if src}
    <img class={`rounded-full object-center object-cover ${sizeClass}`} {src} {alt} />
  {:else}
    <Icon class={`rounded-full text-cyan-500 ${sizeClass}`} icon="hugeicons:user-circle-02" />
  {/if}
</AvatarLayout>
