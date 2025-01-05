<script lang="ts">
    type HomeReviewsSectionItemTextProps = {
        content: string;
    };

    const {content}: HomeReviewsSectionItemTextProps = $props();
</script>

<p class="text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary mb-2 line-clamp-5">"{content}"</p>