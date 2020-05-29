/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import React from 'react';
import { Box, Link, LinkProps } from 'pouncejs';
import { LinkifyProps } from 'linkifyjs/react';

const OriginalReactLinkify = React.lazy(() =>
  import(/* webpackChunkName: "linkify" */ 'linkifyjs/react.js')
) as React.FC<LinkifyProps>;

const linkifyOptions = {
  defaultProtocol: 'https',
  tagName: () => Link,
  attributes: {
    color: 'blue300' as const,
    external: true,
  } as LinkProps,
};

const Linkify: React.FC = ({ children }) => {
  return (
    <Box wordBreak="break-word" fontSize="medium">
      <React.Suspense fallback={<div>{children}</div>}>
        <OriginalReactLinkify options={linkifyOptions as any}>{children}</OriginalReactLinkify>
      </React.Suspense>
    </Box>
  );
};

export default Linkify;
