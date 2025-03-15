<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps'

  type DropdownTriggerProps = LayoutProps & {
    shown: boolean
    triggerRef: HTMLDivElement | undefined
  }

  let { shown = $bindable(), triggerRef = $bindable(), children }: DropdownTriggerProps = $props()
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div bind:this={triggerRef} class="cursor-pointer" onclick={() => (shown = !shown)}>
  {@render children()}
</div>
