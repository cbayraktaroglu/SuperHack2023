<script lang="ts">
	export let triggerFunction: any = '';
	export let buttonText: String = '';

	let initialColor: string = 'black';
	let isHovered: boolean = false;
	let hue: number = 0;
	let buttonStyle: string = `background-color: ${initialColor}`;

	function startHover(): void {
		isHovered = true;
		changeColor();
	}

	function endHover(): void {
		isHovered = false;
		buttonStyle = `background-color: ${initialColor}`;
	}

	function changeColor(): void {
		if (isHovered) {
			hue += 10;
			hue = hue % 360;

			const newColor: string = `hsl(${hue}, 90%, 80%)`;
			buttonStyle = `background-color: ${newColor}`;

			setTimeout(changeColor, 100); // Continuously change color while hovering
		}
	}
</script>

<button
	class="color-button"
	style={buttonStyle}
	on:mouseenter={startHover}
	on:mouseleave={endHover}
	on:click={triggerFunction}
>
	{buttonText}
</button>

<style>
	.color-button {
		padding: 10px 20px;
		border: none;
		cursor: pointer;
		transition: background-color 0.3s ease-in-out;
		color: white;
		border-radius: 15px;
	}
</style>
