<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";

    const {minPrice}: UserService = $props();
</script>

<p class="text-base font-medium">From ${minPrice}</p>
