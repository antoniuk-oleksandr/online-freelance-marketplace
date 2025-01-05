<script lang="ts">
    type HomeReviewsSectionItemNameProps = {
        firstName: string,
        surname: string,
    };

    const {firstName, surname}: HomeReviewsSectionItemNameProps = $props();
</script>

<h3 class="text-xl font-semibold text-cyan-500 dark:text-cyan-500">{firstName} {surname}</h3>
