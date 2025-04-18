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
    class="w-[90%] max-w-xl sm:w-[70%] lg:w-[40%] h-auto max-h-[75vh]"
    transition:flyFade={{ y: 10, duration: 300 }}
  >
    <PaperElement styles="!p-0 text-base size-full shadow-lg">
      {@render children()}
    </PaperElement>
  </div>
</ModalBackdrop>
