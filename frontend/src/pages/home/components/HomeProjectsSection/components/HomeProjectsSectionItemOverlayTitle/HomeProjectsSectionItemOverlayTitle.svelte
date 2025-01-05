<script lang="ts">
    import  type {BestService} from "@/types/BestService.ts";

    const {title}: BestService = $props();
</script>

<h3 class="text-xl font-medium text-light-palette-text-primary dark:text-dark-palette-text-primary">
    {title}
</h3>