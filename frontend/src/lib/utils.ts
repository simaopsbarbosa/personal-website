/**
 * expected input: "YYYY-MM-DD HH:MM:SS"
 * output: "Jun 20, 2026"
 */
export function formatDate(dateStr: string): string {
	if (!dateStr) return '';
	
	const date = new Date(dateStr.replace(' ', 'T'));
	
	if (isNaN(date.getTime())) {
		return dateStr;
	}
	
	return date.toLocaleDateString('en-US', {
		day: 'numeric',
		month: 'short',
		year: 'numeric'
	});
}
