import React from 'react';

function CompanyCard({ company }) {
  const changeValue = company.today_points.slice(-1)[0].change_value
  const changeRatio = company.today_points.slice(-1)[0].change_ratio
  const openValue = company.today_points.slice(-1)[0].open
  const closeValue = company.today_points.slice(-1)[0].close
  const tradeValue = company.today_points.slice(-1)[0].trade_value
  const tradeVolume = company.today_points.slice(-1)[0].trade_volume
  const tradeDate = company.today_points.slice(-1)[0].trade_date
  return (
    <div className="bg-white p-4 rounded-md shadow-md mb-4 font-rubik" dir="rtl">
      <h2 className="text-xl font-semibold mb-2">{company.company_name}</h2>
      <p className="mb-1">{company.short_name} <span className="bg-green-100 text-green-800 text-xs font-extrabold px-2.5 py-0.5 rounded-full dark:bg-green-900 dark:text-green-300"> رمز {company.symbol_code}</span></p>
      <p className='text-sm font-semibold mb-2'>نقاط اليوم 💰</p>
      <p className='text-sm font-semibold mb-2'>تاريخ: {tradeDate}</p>
      <div >
        {changeValue > 0 ? <span className="bg-blue-100 text-blue-800 text-xs m-1 px-2.5 py-0.5 rounded-full dark:bg-blue-900 dark:text-blue-300 font-extrabold">ربح {changeValue}%</span> : <span className="bg-red-100 text-red-800 text-xs font-extrabold px-2.5 py-0.5 rounded-full dark:bg-red-900 dark:text-red-300">خسارة {changeValue}%</span>}
        <span className="bg-amber-100 text-amber-800 text-xs m-1 px-2.5 py-0.5 rounded-full dark:bg-amber-900 dark:text-amber-300 font-extrabold">نسبة التغير {changeRatio}%</span>
        <span className="bg-zinc-100 text-zinc-800 text-xs px-2.5 m-1 py-0.5 rounded-full dark:bg-zinc-900 dark:text-zinc-300 font-extrabold">افتتاح {openValue} ريال</span>
        <span className="bg-zinc-100 text-zinc-800 text-xs px-2.5 m-1 py-0.5 rounded-full dark:bg-zinc-900 dark:text-zinc-300 font-extrabold">أغلاق {closeValue} ريال</span>
      </div>
      <div>
        <span className="bg-zinc-100 text-zinc-800 text-xs px-2.5 m-1 py-0.5 rounded-full dark:bg-zinc-900 dark:text-zinc-300 font-extrabold">قيمة التداول {tradeValue} ريال</span>
        <span className="bg-zinc-100 text-zinc-800 text-xs px-2.5 m-1 py-0.5 rounded-full dark:bg-zinc-900 dark:text-zinc-300 font-extrabold">حجم التداول {tradeVolume} سهم</span>
      </div>
    </div>
  );
}

export default CompanyCard;
