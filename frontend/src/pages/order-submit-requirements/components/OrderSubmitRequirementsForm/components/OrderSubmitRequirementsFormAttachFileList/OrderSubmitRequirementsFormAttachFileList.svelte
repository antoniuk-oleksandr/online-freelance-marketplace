<script lang="ts">
  import OrderSubmitRequirementsFormAttachFile from '../OrderSubmitRequirementsFormAttachFile/OrderSubmitRequirementsFormAttachFile.svelte'
  import OrderSubmitRequirementsFormAttachFileListLayout from './OrderSubmitRequirementsFormAttachFileListLayout.svelte'

  type OrderSubmitRequirementsFormAttachFileListProps = {
    files: File[]
    removeFile: (index: number) => void
  }

  const { files, ...rest }: OrderSubmitRequirementsFormAttachFileListProps = $props()
</script>

<OrderSubmitRequirementsFormAttachFileListLayout>
  {#each files as file, index}
    <OrderSubmitRequirementsFormAttachFile {...rest} {index} {file} />
  {/each}
</OrderSubmitRequirementsFormAttachFileListLayout>
