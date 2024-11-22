<script lang="ts">
    import type {ToastElementStore} from "@/types/ToastElementStore.ts";

    const {message}: ToastElementStore = $props();
</script>

<p class="text-base font-medium text-light-palette-text-primary dark:text-dark-palette-text-primary">
    {message}
</p>