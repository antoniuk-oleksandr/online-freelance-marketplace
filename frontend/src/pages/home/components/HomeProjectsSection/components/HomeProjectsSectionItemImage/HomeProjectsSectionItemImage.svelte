<script lang="ts">
    import type {BestService} from "@/types/BestService.ts";

    const {title, image}: BestService = $props();
</script>

<img src={image} alt={title} class="w-full h-128 object-center rounded-xl object-cover"/>
