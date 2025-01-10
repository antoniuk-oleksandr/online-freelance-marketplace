<script lang="ts">
    type TitleProps = {
        text: string;
    };

    const {text}: TitleProps = $props();
</script>

<h2 class="text-xl capitalize font-semibold">{text}</h2>
