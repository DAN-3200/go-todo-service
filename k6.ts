import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
	vus: 100, // Número de usuários virtuais
	iterations: 100, // Número de iterações por usuário
};

export default function () {
	const randomId = Math.floor(Math.random() * 1000);
	http.post(
		'http://localhost:8200/todo',
		JSON.stringify({
			title: `[${randomId}] Teste de carga `,
			content: `[${randomId}] Descrição do teste de carga `,
		}),
		{
			headers: { 'Content-Type': 'application/json' },
		},
	);

	sleep(1); // Pausa de 1 segundo entre as requisições
}
