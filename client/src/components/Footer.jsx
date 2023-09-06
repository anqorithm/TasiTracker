import React from 'react';

function Footer() {
    return (
        <div className="bg-gray-200 p-6">
            <div className="container mx-auto text-center md:text-left">
                <div className="flex flex-wrap justify-center md:justify-between">
                    <div className="mb-4 md:mb-0">
                        <h6 className="text-gray-700 font-bold mb-2">Company Name</h6>
                        <ul className="list-none mb-0 space-y-2">
                            <li><a href="#" className="text-gray-600 hover:text-gray-900">About Us</a></li>
                            <li><a href="#" className="text-gray-600 hover:text-gray-900">Services</a></li>
                            <li><a href="#" className="text-gray-600 hover:text-gray-900">Contact</a></li>
                        </ul>
                    </div>
                    <div className="mb-4 md:mb-0">
                        <h6 className="text-gray-700 font-bold mb-2">Support</h6>
                        <ul className="list-none mb-0 space-y-2">
                            <li><a href="#" className="text-gray-600 hover:text-gray-900">FAQ</a></li>
                            <li><a href="#" className="text-gray-600 hover:text-gray-900">Help Center</a></li>
                            <li><a href="#" className="text-gray-600 hover:text-gray-900">Privacy Policy</a></li>
                        </ul>
                    </div>
                </div>
                <hr className="my-4 border-gray-300 w-full"/>
                <div className="text-gray-600 text-sm">
                    &copy; {(new Date().getFullYear())} Company Name. All rights reserved.
                </div>
            </div>
        </div>
    );
}
 export default Footer;
