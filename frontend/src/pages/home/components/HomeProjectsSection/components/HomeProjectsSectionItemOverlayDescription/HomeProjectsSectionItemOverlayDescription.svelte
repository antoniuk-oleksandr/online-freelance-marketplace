<script lang="ts">
    import type {BestService} from "@/types/BestService.ts";

    const {description}: BestService = $props();
</script>

<p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary line-clamp-2 h-12">
    {description}
</p>