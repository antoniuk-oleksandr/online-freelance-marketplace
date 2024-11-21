<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
    import Icon from "@iconify/svelte";

    type ErrorComponentProps = {
        error: ResponseErrorEnum,
    };

    const {error}: ErrorComponentProps = $props();
</script>

<PaperElement styles="flex text-base !border-0 md:!border !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block flex-col items-center gap-4 p-8 max-w-168">
    <Icon icon="material-symbols:error" class="text-red-500" width="64" height="64"/>
    <h1 class="text-2xl font-bold">Error: {error}</h1>
    <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        We couldn't process your request due to the following error:
        <strong>{error}</strong>.
    </p>
    <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        Please try again later. If the issue persists, contact support for assistance.
    </p>
</PaperElement>

