<script lang="ts">
    type SuccessMessageSubtextProps = {
        subText: string,
    }

    const {subText}: SuccessMessageSubtextProps = $props();
</script>

<p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {subText}
</p>