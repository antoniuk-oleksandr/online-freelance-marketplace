<script lang="ts">
  import ModalBackdrop from '@/common-components/Modal/components/ModalBackdrop/ModalBackdrop.svelte'
  import type { LayoutProps } from '@/types/LayoutProps.ts'
  import PaperElement from '@/common-components/PaperElement/PaperElement.svelte'
  import type { ModalStore } from '@/types/ModalStore.ts'
  import { flyFade } from '@/utils/utils'

  type ModalLayoutProps = LayoutProps & {
    modalData: ModalStore
  }

  let modalContentRef = $state<HTMLElement | undefined>()

  const { children, modalData }: ModalLayoutProps = $props()
</script>

<ModalBackdrop {modalContentRef} {modalData}>
  <div
    bind:this={modalContentRef}
    class="w-full max-w-[min(100%,45rem)] h-fit max-h-[100svh-3rem] md:max-h-[75vh] grid place-items-center"
    transition:flyFade={{ y: 10, duration: 300 }}
  >
    <PaperElement styles="!p-0 text-base size-full !max-h-[calc(100svh-3rem)] shadow-lg overflow-y-auto">
      {@render children()}
    </PaperElement>
  </div>
</ModalBackdrop>
