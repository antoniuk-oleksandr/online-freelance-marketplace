<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'

  type LabeledInputElementProps = LayoutProps & {
    styles?: string
  }

  const { children, styles }: LabeledInputElementProps = $props()
</script>

<div class="{styles} w-full font-medium flex flex-col gap-1">
  {@render children()}
</div>
