<script lang="ts">
    import Icon from "@iconify/svelte";
    import type {BestService} from "@/types/BestService.ts";

    const {icon}: BestService = $props();
</script>

<Icon
        icon={icon ?? "mdi:file-document-outline"}
        class="text-cyan-500 dark:text-cyan-500" width="32"
        height="32"
/>