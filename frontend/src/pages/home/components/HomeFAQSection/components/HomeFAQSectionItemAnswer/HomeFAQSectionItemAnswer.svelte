<script lang="ts">
    type HomeFAQSectionItemAnswerProps = {
        answer: string,
    };

    const { answer }:  HomeFAQSectionItemAnswerProps = $props();
</script>

<p class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {answer}
</p>