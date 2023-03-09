/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * Monta is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Monta is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Monta.  If not, see <https://www.gnu.org/licenses/>.
 */

/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable jsx-a11y/anchor-is-valid */
import React from 'react'
import {Demos} from './Demos'

export function DemosDrawer() {
  return (
    <>
      {/* begin::Demos drawer */}
      <div
        id='mt_engage_demos'
        className='bg-body'
        data-mt-drawer='true'
        data-mt-drawer-name='explore'
        data-mt-drawer-activate='true'
        data-mt-drawer-overlay='true'
        data-mt-drawer-width="{default:'350px', 'lg': '475px'}"
        data-mt-drawer-direction='end'
        data-mt-drawer-toggle='#mt_engage_demos_toggle'
        data-mt-drawer-close='#mt_engage_demos_close'
      >
        {/* begin::Card  */}
        <div className='card shadow-none w-100'>
          {/* begin::Header */}
          <div className='card-header' id='mt_explore_header'>
            <h5 className='card-title fw-bold text-gray-600'>
              Explore {process.env.NEXT_PUBLIC_APP_THEME_NAME}
            </h5>

            <div className='card-toolbar'>
              <button
                type='button'
                className='btn btn-sm btn-icon explore-btn-dismiss me-n5'
                id='mt_engage_demos_close'
              >
                <img src='/media/icons/duotune/arrows/arr061.svg' className='svg-icon-2' />
              </button>
            </div>
          </div>
          {/* end::Header */}
          {/* begin::Body */}
          <div className='card-body' id='mt_explore_body'>
            {/* begin::Content */}
            <div
              id='mt_explore_scroll'
              className='scroll-y me-n5 pe-5'
              data-mt-scroll='true'
              data-mt-scroll-height='auto'
              data-mt-scroll-wrappers='#mt_engage_demos_body'
              data-mt-scroll-dependencies='#mt_engage_demos_header'
              data-mt-scroll-offset='5px'
            >
              <div className='mb-7'>
                <div className='d-flex flex-stack'>
                  <h3 className='mb-0'>{process.env.NEXT_PUBLIC_APP_THEME_NAME} Licenses</h3>

                  <a href='https://themeforest.net/licenses/standard' className='fw-bold'>
                    License FAQs
                  </a>
                </div>
              </div>

              <div className='rounded border border-dashed border-gray-300 py-4 px-6 mb-5'>
                <div className='d-flex flex-stack'>
                  <div className='d-flex flex-column'>
                    <div className='d-flex align-items-center mb-1'>
                      <div className='fs-6 fw-bold text-gray-800 fw-bold mb-0 me-1'>
                        Regular License
                      </div>

                      <i
                        className='text-gray-400 fas fa-exclamation-circle ms-1 fs-7'
                        data-bs-toggle='popover'
                        data-bs-trigger='hover'
                        data-bs-placement='top'
                        data-bs-content='Use, by you or one client in a single end product which end users are not charged for.'
                      ></i>
                    </div>
                    <div className='fs-7 text-muted'>
                      For single end product used by you or one client
                    </div>
                  </div>

                  <div className='text-nowrap'>
                    <span className='text-muted fs-7 fw-bold'>$</span>
                    <span className='text-dark fs-1 fw-bolder'>39</span>
                  </div>
                </div>
              </div>

              <div className='rounded border border-dashed border-gray-300 py-4 px-6 mb-5'>
                <div className='d-flex flex-stack'>
                  <div className='d-flex flex-column'>
                    <div className='d-flex align-items-center mb-1'>
                      <div className='fs-6 fw-bold text-gray-800 fw-bold mb-0 me-1'>
                        Extended License
                      </div>

                      <i
                        className='text-gray-400 fas fa-exclamation-circle ms-1 fs-7'
                        data-bs-toggle='popover'
                        data-bs-trigger='hover'
                        data-bs-placement='top'
                        data-bs-content='Use, by you or one client, in a single end product which end users can be charged for.'
                      ></i>
                    </div>
                    <div className='fs-7 text-muted'>For single end product with paying users.</div>
                  </div>

                  <div className='text-nowrap'>
                    <span className='text-muted fs-7 fw-bold'>$</span>
                    <span className='text-dark fs-1 fw-bolder'>939</span>
                  </div>
                </div>
              </div>

              <div className='rounded border border-dashed border-gray-300 py-4 px-6 mb-5'>
                <div className='d-flex flex-stack'>
                  <div className='d-flex flex-column'>
                    <div className='d-flex align-items-center mb-1'>
                      <div className='fs-6 fw-bold text-gray-800 fw-bold mb-0 me-1'>
                        Custom License
                      </div>
                    </div>
                    <div className='fs-7 text-muted'>Reach us for custom license offers.</div>
                  </div>

                  <div className='text-nowrap'>
                    <a href='https://keenthemes.com/contact/' className='btn btn-sm btn-success'>
                      Contact Us
                    </a>
                  </div>
                </div>
              </div>

              <a href={process.env.NEXT_APP_PURCHASE_URL} className='btn btn-primary mb-15 w-100'>
                Buy Now
              </a>
              <Demos />
            </div>
            {/* end::Content */}
          </div>
          {/* end::Body */}
        </div>
        {/* begin::Card */}
      </div>
      {/* end::Demos drawer */}
    </>
  )
}
