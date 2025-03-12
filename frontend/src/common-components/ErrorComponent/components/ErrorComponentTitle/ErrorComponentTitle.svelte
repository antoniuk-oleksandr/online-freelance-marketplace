<script lang="ts">
    import type {ResponseErrorEnum} from "@/types/ResponseErrorEnum";

    type ErrorComponentTitleProps = {
        error: ResponseErrorEnum | undefined,
    }

    const {error}: ErrorComponentTitleProps = $props();
</script>

<h1 class="text-2xl font-bold text-center">Error: {error}</h1>
