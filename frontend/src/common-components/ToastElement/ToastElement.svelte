<script lang="ts">
    import ToastElementLayout from "@/common-components/ToastElement/ToastElementLayout.svelte";
    import {toastElementStore} from "@/common-components/ToastElement/store/toast-element-store.ts";
    import type {ToastElementStore} from "@/types/ToastElementStore.ts";
    import ToastElementIcon from "@/common-components/ToastElement/components/ToastElementIcon/ToastElementIcon.svelte";
    import ToastElementMessage
        from "@/common-components/ToastElement/components/ToastElementMessage/ToastElementMessage.svelte";
    import ToastElementCloseButton
        from "@/common-components/ToastElement/components/ToastElementCloseButton/ToastElementCloseButton.svelte";
    import {showOpenToastAnimation} from "@/common-components/ToastElement/helpers.ts";

    let toastValue = $state<ToastElementStore | undefined>();
    toastElementStore.subscribe((value) => {
        toastValue = value;
    });

    $effect(() => {
        if (toastValue && toastValue.show) {
            const timeout = showOpenToastAnimation();
            return () => clearTimeout(timeout);
        }
    })
</script>

{#if toastValue && toastValue.show}
    <ToastElementLayout {...toastValue}>
        <ToastElementIcon {...toastValue}/>
        <ToastElementMessage {...toastValue}/>
        <ToastElementCloseButton/>
    </ToastElementLayout>
{/if}
