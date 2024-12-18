<script lang="ts">
    import ModalBackdrop from "@/common-components/Modal/components/ModalBackdrop/ModalBackdrop.svelte";
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import type {ModalStore} from "@/types/ModalStore.ts";

    type ModalLayoutProps = LayoutProps & {
        modalData: ModalStore;
    }

    const {children, modalData}: ModalLayoutProps = $props();
</script>

<ModalBackdrop modalData={modalData}>
    <PaperElement styles="!p-0 text-base w-full w-[90%] max-w-xl sm:w-[70%] lg:w-[40%] h-auto max-h-[75vh] shadow-lg">
        {@render children()}
    </PaperElement>
</ModalBackdrop>