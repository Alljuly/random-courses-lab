import { X } from "lucide-react";

export function Modal() {
  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="relative bg-[#F0F2F5] rounded-lg p-10 w-full max-w-[500px] shadow-lg">
        {/* Botão fechar */}
        <button className="absolute right-6 top-6 text-[#969CB2] hover:text-black transition-colors">
          <X size={20} />
        </button>
        <h2 className="text-2xl font-bold text-gray-900 mb-8">
          Cadastrar transação
        </h2>

        <form className="flex flex-col gap-4">
          <input
            type="text"
            placeholder="Nome"
            className="bg-white rounded-md p-4 text-gray-900 border border-transparent focus:border-green-500 outline-none"
          />
          <input
            type="number"
            placeholder="Preço"
            className="bg-white rounded-md p-4 text-gray-900 border border-transparent focus:border-green-500 outline-none"
          />
          <div className="flex gap-2">
            <button
              type="button"
              className="flex-1 flex items-center justify-center gap-2 border border-[#d7d7d7] rounded-md p-4 bg-white text-[#33CC95] font-semibold hover:border-green-500 transition-colors"
            >
              <svg width="20" height="20" fill="none" viewBox="0 0 24 24">
                <path
                  fill="#33CC95"
                  d="M12 3v18m9-9H3"
                  stroke="#33CC95"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
              Entrada
            </button>
            <button
              type="button"
              className="flex-1 flex items-center justify-center gap-2 border border-[#d7d7d7] rounded-md p-4 bg-white text-[#E52E4D] font-semibold hover:border-red-500 transition-colors"
            >
              <svg width="20" height="20" fill="none" viewBox="0 0 24 24">
                <path
                  fill="#E52E4D"
                  d="M5 12h14"
                  stroke="#E52E4D"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
              Saída
            </button>
          </div>
          <input
            type="text"
            placeholder="Categoria"
            className="bg-white rounded-md p-4 text-gray-900 border border-transparent focus:border-green-500 outline-none"
          />
          <button
            type="submit"
            className="bg-[#33CC95] text-white rounded-md p-4 font-bold mt-2 hover:bg-green-600 transition-colors w-full"
          >
            Cadastrar
          </button>
        </form>
      </div>
    </div>
  );
}
