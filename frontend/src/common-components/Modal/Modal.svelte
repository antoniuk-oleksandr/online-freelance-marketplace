<script lang="ts">
  import ModalLayout from '@/common-components/Modal/ModalLayout.svelte'
  import ModalHeader from '@/common-components/Modal/components/ModalHeader/ModalHeader.svelte'
  import { modalStore } from '@/common-stores/modal-store'
  import type { ModalStore } from '@/types/ModalStore.ts'
  import { handleModalOpen } from '@/common-components/Modal/handlers'

  let modalData = $state<ModalStore | undefined>()
  modalStore.subscribe((value) => (modalData = value))

  $effect(() => handleModalOpen(modalData))
</script>

{#if modalData && modalData.renderContent && modalData.isOpened}
  <ModalLayout {modalData}>
    <ModalHeader headerStyles={modalData.headerStyles} title={modalData.title} />
    {@render modalData.renderContent(modalData.props)}
  </ModalLayout>
{/if}
