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

const BIRTH_DATE = new Date('2005-12-22T13:30:00Z');

/**
 * calculates the exact age from my birth date
 * up to the given datetime input.
 */
export function calculateAge(dateInput: string | Date | number = new Date()) {
	const normalizedInput = typeof dateInput === 'string' ? dateInput.replace(' ', 'T') : dateInput;
	const now = new Date(normalizedInput);

	if (isNaN(now.getTime()) || now.getTime() < BIRTH_DATE.getTime()) {
		return {
			years: 0,
			days: 0,
			hours: 0,
			minutes: 0,
			seconds: 0,
			decimal: 0,
			formatted: '0y 0d 0h 0m 0s'
		};
	}

	let years = now.getFullYear() - BIRTH_DATE.getFullYear();
	const birthDayThisYear = new Date(BIRTH_DATE);
	birthDayThisYear.setFullYear(now.getFullYear());

	if (now.getTime() < birthDayThisYear.getTime()) {
		years--;
		birthDayThisYear.setFullYear(now.getFullYear() - 1);
	}

	const msSinceBirthday = now.getTime() - birthDayThisYear.getTime();
	const secondsSinceBirthday = Math.floor(msSinceBirthday / 1000);

	const days = Math.floor(secondsSinceBirthday / (24 * 3600));
	let tempSec = secondsSinceBirthday % (24 * 3600);
	const hours = Math.floor(tempSec / 3600);
	tempSec %= 3600;
	const minutes = Math.floor(tempSec / 60);
	const seconds = tempSec % 60;

	const nextBirthday = new Date(birthDayThisYear);
	nextBirthday.setFullYear(birthDayThisYear.getFullYear() + 1);
	const totalMsInYear = nextBirthday.getTime() - birthDayThisYear.getTime();
	const decimal = years + (now.getTime() - birthDayThisYear.getTime()) / totalMsInYear;

	return {
		years,
		days,
		hours,
		minutes,
		seconds,
		decimal,
		formatted: `${years}y ${days}d ${hours}h ${minutes}m ${seconds}s`
	};
}
