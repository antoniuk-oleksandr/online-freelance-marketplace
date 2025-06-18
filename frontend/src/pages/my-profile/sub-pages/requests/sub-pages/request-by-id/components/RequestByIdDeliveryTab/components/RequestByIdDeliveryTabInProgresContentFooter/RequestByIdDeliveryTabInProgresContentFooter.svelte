<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import DividerElement from '@/common-components/DividerElement/DividerElement.svelte'
  import Icon from '@iconify/svelte'
  import RequestByIdDeliveryTabInProgresContentFooterLayout from './RequestByIdDeliveryTabInProgresContentFooterLayout.svelte'
  import { getContext } from 'svelte'
  import type { Writable } from 'svelte/store'

  type RequestByIdDeliveryTabInProgresContentFooterProps = {
    files: File[]
  }

  const { files }: RequestByIdDeliveryTabInProgresContentFooterProps = $props()
  const formIsSubmittingStore = getContext<Writable<boolean>>('formIsSubmittingStore')
</script>

<RequestByIdDeliveryTabInProgresContentFooterLayout>
  <DividerElement />
  <div class="flex flex-col md:flex-row md:justify-between md:items-center gap-3">
    <div class="text-sm text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
      {#if files.length > 0}
        <div class="flex items-center gap-1">
          <Icon icon="hugeicons:file-02" width="20" height="20" />
          <span>{files.length} file{files.length !== 1 ? 's' : ''} attached</span>
        </div>
      {:else}
        <div class="flex items-center gap-1">
          <Icon icon="hugeicons:alert-02" width="20" height="20" />
          <span>No files attached yet</span>
        </div>
      {/if}
    </div>
    <Button type="submit" loading={$formIsSubmittingStore} styles="!gap-2 !min-w-40" color="green">
      <Icon icon="hugeicons:checkmark-circle-02" width="20" height="20" />
      <span>Deliver Order</span>
    </Button>
  </div>
</RequestByIdDeliveryTabInProgresContentFooterLayout>
