<script lang="ts">
    import type {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

    type ErrorComponentFirstLineTextProps = {
        error: ResponseErrorEnum | undefined
    }

    const {error}:ErrorComponentFirstLineTextProps = $props();
</script>

<p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    We couldn't process your request due to the following error:
    <strong>{error}</strong>.
</p>