<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import Icon from "@iconify/svelte";
    import {errorStore} from "@/common-stores/error-store.ts";
    import type {ErrorStore} from "@/types/ErrorStore.ts";
    import ErrorComponentLayout from "@/common-components/ErrorComponent/ErrorComponentLayout.svelte";

    let errorData = $state<ErrorStore>();
    errorStore.subscribe((value) => errorData = value);

    $effect(() => {
        if (errorData && errorData.shown) document.body.style.overflow = "hidden";
        else document.body.style.overflow = "auto";
    })
</script>

{#if errorData && errorData.shown}
    <ErrorComponentLayout>
        <PaperElement
                styles="flex text-base !border-0 md:!border !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block flex-col items-center gap-4 p-6 max-w-168">
            <Icon icon="material-symbols:error" class="text-red-500" width="64" height="64"/>
            <h1 class="text-2xl font-bold">Error: {errorData.error}</h1>
            <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
                We couldn't process your request due to the following error:
                <strong>{errorData.error}</strong>.
            </p>
            <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
                Please try again later. If the issue persists, contact support for assistance.
            </p>
        </PaperElement>
    </ErrorComponentLayout>
{/if}
