<script lang="ts">
  import { onMount, onDestroy } from 'svelte'
  import OrderSubmitRequirementsFormAttachFile from '../OrderSubmitRequirementsFormAttachFile/OrderSubmitRequirementsFormAttachFile.svelte'
  import OrderSubmitRequirementsFormAttachFileListLayout from './OrderSubmitRequirementsFormAttachFileListLayout.svelte'

  type OrderSubmitRequirementsFormAttachFileListProps = {
    files: File[]
    removeFile: (index: number) => void
    styles?: string
  }

  let isMobile: boolean = $state(window.innerWidth <= 768)

  const handleResize = () => (isMobile = window.innerWidth <= 768)
  onMount(() => window.addEventListener('resize', handleResize))
  onDestroy(() => window.removeEventListener('resize', handleResize))

  const { files, ...rest }: OrderSubmitRequirementsFormAttachFileListProps = $props()
</script>

<OrderSubmitRequirementsFormAttachFileListLayout {...rest} {isMobile}>
  {#each files as file, index}
    <OrderSubmitRequirementsFormAttachFile {...rest} {isMobile} {index} {file} />
  {/each}
</OrderSubmitRequirementsFormAttachFileListLayout>
