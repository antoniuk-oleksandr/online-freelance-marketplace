<script lang="ts">
  import Icon from '@iconify/svelte'
  import { handleDownloadButtonClick } from '../../handlers'
  import OrderByIdDeliveryTabCompleteContentFileLayout from './OrderByIdDeliveryTabCompleteContentFileLayout.svelte'

  type OrderByIdDeliveryTabCompleteContentFileProps = {
    file: string
  }

  const { file }: OrderByIdDeliveryTabCompleteContentFileProps = $props()
</script>

<OrderByIdDeliveryTabCompleteContentFileLayout>
  <div
    class="bg-light-palette-action-hover dark:bg-dark-palette-action-hover rounded-lg p-2 flex items-center justify-center"
  >
    <Icon
      icon="hugeicons:file-02"
      width="24"
      height="24"
      class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
    />
  </div>
  <div class="min-w-0 mt-2 flex-1">
    <p
      class="text-base font-medium text-light-palette-text-primary dark:text-dark-palette-text-primary truncate"
    >
      {file.split('/').pop()}
    </p>
  </div>
  <button
    onclick={() => handleDownloadButtonClick(file)}
    class="text-green-600 dark:text-green-500 hover:text-green-700 dark:hover-text-green-300 dark:hover:bg-green-900 duration-200 p-1 rounded-full hover:bg-green-50 transition-colors"
  >
    <Icon icon="hugeicons:download-04" width="24" height="24" />
  </button>
</OrderByIdDeliveryTabCompleteContentFileLayout>
