<script lang="ts">
    type SpinnerProps = {
        size?: string,
        color?: string
    };

    const {size = "size-14", color = 'border-l-cyan-500'}: SpinnerProps = $props();
</script>


<div class="animate-spin border-2 border-transparent rounded-full {size} {color}"></div>