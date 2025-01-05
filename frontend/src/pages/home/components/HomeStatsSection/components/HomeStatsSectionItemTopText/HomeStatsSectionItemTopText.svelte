<script lang="ts">
    type StatsSectionItemTopTextProps = {
        topText: string,
    };

    const {topText}: StatsSectionItemTopTextProps = $props();
</script>

<div class="text-4xl font-bold text-cyan-500 dark:text-cyan-500">{topText}</div>
