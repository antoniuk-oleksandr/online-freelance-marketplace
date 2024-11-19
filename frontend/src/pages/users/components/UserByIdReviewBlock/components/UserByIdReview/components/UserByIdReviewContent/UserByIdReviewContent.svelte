<script lang="ts">
    import type {Review} from "@/types/Review.ts";

    const {content}: Review = $props();
</script>

<div
        class="text-base text-light-palette-text-secondary dark:text-dark-palette-text-secondary"
>
    <p>{content}</p>
</div>