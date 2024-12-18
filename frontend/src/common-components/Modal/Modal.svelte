<script lang="ts">
    import ModalLayout from "@/common-components/Modal/ModalLayout.svelte";
    import ModalHeader from "@/common-components/Modal/components/ModalHeader/ModalHeader.svelte";
    import {modalStore} from "@/common-stores/modal-store.ts";
    import type {ModalStore} from "@/types/ModalStore.ts";
    import {handleModalOpen} from "@/common-components/Modal/handlers.ts";

    let modalData = $state<ModalStore | undefined>();
    modalStore.subscribe((value) => modalData = value);

    let exitAnimation = $state(false);

    $effect(() => {
        let timeout: number;

        if (modalData && modalData.isOpened) exitAnimation = true;
        else timeout = setTimeout(() => exitAnimation = false, 300);

        return () => clearTimeout(timeout);
    })

    let contentElement = $state<HTMLElement | undefined>();

    $effect(() => handleModalOpen(modalData));

    $effect(() => {
        if (modalData && modalData.renderContent && contentElement && modalData.isOpened) {
            modalData.renderContent(contentElement);
        }
    })
</script>

{#if modalData && exitAnimation && modalData.renderContent}
    <ModalLayout modalData={modalData}>
        <ModalHeader headerStyles={modalData.headerStyles} title={modalData.title}/>
        <div class="" bind:this={contentElement}></div>
    </ModalLayout>
{/if}
