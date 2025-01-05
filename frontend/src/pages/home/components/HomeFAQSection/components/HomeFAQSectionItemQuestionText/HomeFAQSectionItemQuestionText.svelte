<script lang="ts">
    type HomeFAQSectionItemQuestionTextProps = {
        question: string;
    };

    const {question}: HomeFAQSectionItemQuestionTextProps = $props();
</script>

<h3 class="text-xl font-medium text-light-palette-text-primary dark:text-dark-palette-text-primary">
    {question}
</h3>