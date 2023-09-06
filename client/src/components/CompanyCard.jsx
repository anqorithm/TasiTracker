import React from 'react';

function CompanyCard({ company }) {
    const changeValue = company.today_points.slice(-1)[0].change_value
  return (
    <div className="bg-white p-4 rounded-md shadow-md mb-4 font-rubik" dir="rtl">
      <h2 className="text-xl font-semibold mb-2">{company.company_name}</h2>
      <p className="mb-1">{company.short_name} <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-green-900 dark:text-green-300">{company.symbol_code}</span></p>
     {changeValue > 0 ? <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-blue-900 dark:text-blue-300">ربح {changeValue}%</span> : <span className="bg-red-100 text-red-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-red-900 dark:text-red-300">خسارة {changeValue}%</span>}
    </div>
  );
}

export default CompanyCard;
