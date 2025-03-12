<script lang="ts">
  import OrderSubmitRequirementsFormAttachFileCloseButtonLayout from './OrderSubmitRequirementsFormAttachFileCloseButtonLayout.svelte'
  import Iconify from '@iconify/svelte'

  type OrderSubmitRequirementsFormAttachFileCloseButtonProps = {
    hovered: boolean
    index: number
    removeFile: (index: number) => void
  }

  const { hovered, ...rest }: OrderSubmitRequirementsFormAttachFileCloseButtonProps = $props()
</script>

{#if hovered}
  <OrderSubmitRequirementsFormAttachFileCloseButtonLayout {...rest}>
    <Iconify width={14} height={14} icon="ci:close-lg" />
  </OrderSubmitRequirementsFormAttachFileCloseButtonLayout>
{/if}
