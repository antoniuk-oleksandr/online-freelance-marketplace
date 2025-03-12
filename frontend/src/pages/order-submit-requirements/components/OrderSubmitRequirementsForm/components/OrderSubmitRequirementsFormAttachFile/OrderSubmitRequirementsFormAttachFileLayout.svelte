<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type OrderSubmitRequirementsFormAttachFileLayoutProps = LayoutProps & {
    hovered: boolean
  }

  let { children, hovered = $bindable() }: OrderSubmitRequirementsFormAttachFileLayoutProps =
    $props()
</script>

<!-- svelte-ignore a11y_mouse_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  onmouseover={() => (hovered = true)}
  onmouseleave={() => (hovered = false)}
  class="rounded-xl !duration-1000 !ease-out relative aspect-square size-full"
>
  {@render children()}
</div>
