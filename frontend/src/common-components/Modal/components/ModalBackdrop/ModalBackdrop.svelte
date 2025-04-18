<script lang="ts">
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import { handleModalBackdropClick } from '@/common-components/Modal/handlers'
  import type { ModalStore } from '@/types/ModalStore.ts'
  import { fade } from 'svelte/transition'

  type ModalBackdropProps = LayoutProps & {
    modalData: ModalStore
    modalContentRef: HTMLElement | undefined
  }

  let modalBackdropRef = $state<HTMLElement | undefined>()
  const { children, modalContentRef }: ModalBackdropProps = $props()
</script>

<div
  transition:fade={{ duration: 300 }}
  bind:this={modalBackdropRef}
  aria-hidden="true"
  onclick={(e) => handleModalBackdropClick(e, modalContentRef, modalBackdropRef)}
  class="bg-black bg-opacity-50 z-50 fixed h-svh w-svw grid place-items-center p-6"
>
  {@render children()}
</div>
