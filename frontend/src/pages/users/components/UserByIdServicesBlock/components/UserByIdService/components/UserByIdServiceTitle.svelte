<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";

    const {title}: UserService = $props();
</script>

<p class="text-lg font-semibold mb-2 line-clamp-2 h-14">{title}</p>