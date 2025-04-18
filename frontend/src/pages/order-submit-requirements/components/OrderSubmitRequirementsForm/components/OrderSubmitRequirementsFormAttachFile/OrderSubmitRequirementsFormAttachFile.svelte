<script lang="ts">
  import OrderSubmitRequirementsFormAttachFileCloseButton from '../OrderSubmitRequirementsFormAttachFileCloseButton/OrderSubmitRequirementsFormAttachFileCloseButton.svelte'
  import OrderSubmitRequirementsFormAttachFileLayout from './OrderSubmitRequirementsFormAttachFileLayout.svelte'
  import OrderSubmitRequirementsFormAttachFileImage from '../OrderSubmitRequirementsFormAttachFileImage/OrderSubmitRequirementsFormAttachFileImage.svelte'
  import OrderSubmitRequirementsFormAttachNormalFile from '../OrderSubmitRequirementsFormAttachNormalFile/OrderSubmitRequirementsFormAttachNormalFile.svelte'

  type OrderSubmitRequirementsFormAttachFileProps = {
    file: File
    index: number
    removeFile: (index: number) => void
    isMobile: boolean,
  }

  const { file, ...rest }: OrderSubmitRequirementsFormAttachFileProps = $props()
  let hovered = $state(false)
</script>

<OrderSubmitRequirementsFormAttachFileLayout bind:hovered>
  <OrderSubmitRequirementsFormAttachFileCloseButton {...rest} {hovered} />
  {#if file.type.includes('image')}
    <OrderSubmitRequirementsFormAttachFileImage {file} />
  {:else}
    <OrderSubmitRequirementsFormAttachNormalFile {file} />
  {/if}
</OrderSubmitRequirementsFormAttachFileLayout>
