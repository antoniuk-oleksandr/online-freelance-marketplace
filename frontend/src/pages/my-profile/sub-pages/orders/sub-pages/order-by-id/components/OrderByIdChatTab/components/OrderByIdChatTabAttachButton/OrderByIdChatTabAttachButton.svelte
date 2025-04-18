<script lang="ts">
  import Button from '@/common-components/Button/Button.svelte'
  import { modalStore } from '@/common-stores/modal-store'
  import Icon from '@iconify/svelte'
  import { OrderByIdChatTabAttachModelContent } from '../OrderByIdChatTabAttachModelContent/OrderByIdChatTabAttachModelContent.svelte'

  const openAttachModal = () => {
    modalStore.update((prev) => ({
      isOpened: true,
      title: 'Upload files',
      renderContent: OrderByIdChatTabAttachModelContent,
    }))
  }
</script>

<Button clickAction={openAttachModal} styles="!rounded-full !px-3 !h-12 !w-12 !p-0" type="submit">
  <Icon icon="hugeicons:attachment-02" width="24" height="24" />
</Button>
