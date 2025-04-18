<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type OrderSubmitRequirementsFormAttachFileListLayoutProps = LayoutProps & {
    isMobile: boolean
    styles?: string
  }

  const { children, isMobile, styles }: OrderSubmitRequirementsFormAttachFileListLayoutProps = $props()
</script>

<div class="{styles} grid {isMobile ? 'grid-cols-2' : 'grid-cols-4'} gap-3">
  {@render children()}
</div>
