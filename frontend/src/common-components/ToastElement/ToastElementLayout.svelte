<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { flyFade } from '@/utils/utils'

  const { children }: LayoutProps = $props()
</script>

<div class="fixed left-6 bottom-6">
  <div
    in:flyFade={{ x: -10, duration: 300 }}
    out:flyFade={{ x: -10, duration: 300 }}
    class="rounded-md ring-1 ring-light-palette-divider dark:ring-dark-palette-divider p-4 bg-light-palette-background-block dark:bg-dark-palette-background-block flex items-center gap-x-3"
  >
    {@render children()}
  </div>
</div>
