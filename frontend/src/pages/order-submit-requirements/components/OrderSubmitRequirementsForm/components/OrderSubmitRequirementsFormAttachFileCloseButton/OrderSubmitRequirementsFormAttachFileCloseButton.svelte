<script lang="ts">
  import OrderSubmitRequirementsFormAttachFileCloseButtonLayout from './OrderSubmitRequirementsFormAttachFileCloseButtonLayout.svelte'
  import Iconify from '@iconify/svelte'

  type OrderSubmitRequirementsFormAttachFileCloseButtonProps = {
    hovered: boolean
    index: number
    removeFile: (index: number) => void
    isMobile: boolean
  }

  const { hovered, isMobile, ...rest }: OrderSubmitRequirementsFormAttachFileCloseButtonProps = $props()
</script>

{#if hovered || isMobile}
  <OrderSubmitRequirementsFormAttachFileCloseButtonLayout {...rest}>
    <Iconify width={isMobile ? 24 : 14} height={isMobile ? 24 : 14} icon="ci:close-lg" />
  </OrderSubmitRequirementsFormAttachFileCloseButtonLayout>
{/if}
