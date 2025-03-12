<script lang="ts">
  import OrderSubmitRequirementsFormAttachNormalFileLayout from './OrderSubmitRequirementsFormAttachNormalFileLayout.svelte'
  import Iconify from '@iconify/svelte'

  type OrderSubmitRequirementsFormAttachNormalFileProps = {
    file: File
  }

  const { file }: OrderSubmitRequirementsFormAttachNormalFileProps = $props()
</script>

<OrderSubmitRequirementsFormAttachNormalFileLayout>
  <Iconify
    width={32}
    height={32}
    icon="hugeicons:file-02"
    class=" text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
  />
  <p class="text-center text-ellipsis overflow-hidden whitespace-nowrap max-w-full">
    {file.name}
  </p>
</OrderSubmitRequirementsFormAttachNormalFileLayout>
