<script lang="ts">
    import type {UserService} from "@/types/UserService.ts";

    type UserByIdServiceTitleProps = UserService & {
        size?: "small" | "normal"
    }

    const {title, size}: UserByIdServiceTitleProps = $props();
</script>

<p class="font-semibold mb-2 line-clamp-2 {size === 'small' ? 'text-base' : 'text-lg'} h-14">{title}</p>