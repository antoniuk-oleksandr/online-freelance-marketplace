<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type AvatarLayoutProps = LayoutProps & {
    styles?: string
  }

  const { children, styles }: AvatarLayoutProps = $props()
</script>

<div class="{styles} size-fit">
  {@render children()}
</div>
