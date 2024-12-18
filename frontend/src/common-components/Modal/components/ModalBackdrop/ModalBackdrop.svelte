<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {handleModalBackdropClick} from "@/common-components/Modal/handlers.ts";
    import type {ModalStore} from "@/types/ModalStore.ts";

    type ModalBackdropProps = LayoutProps & {
        modalData: ModalStore;
    }

    const {children, modalData} : ModalBackdropProps = $props();
</script>

<div
        id="modal-backdrop"
        aria-hidden="true"
        onclick={(e) => handleModalBackdropClick(e)}
        class="bg-black {modalData.isOpened ? 'animate-[fade-in_300ms]' : 'animate-[fade-out_300ms]'} bg-opacity-50 z-50 fixed h-svh w-svw grid place-items-center p-6">
    {@render children()}
</div>